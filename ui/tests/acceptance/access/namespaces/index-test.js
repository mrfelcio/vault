import { currentRouteName } from '@ember/test-helpers';
import { module, test } from 'qunit';
import { setupApplicationTest } from 'ember-qunit';
import { setupMirage } from 'ember-cli-mirage/test-support';
import page from 'vault/tests/pages/access/namespaces/index';
import authPage from 'vault/tests/pages/auth';
import logout from 'vault/tests/pages/logout';

module('Acceptance | Enterprise | /access/namespaces', function(hooks) {
  setupApplicationTest(hooks);
  setupMirage(hooks);

  hooks.beforeEach(function() {
    return authPage.login();
  });

  hooks.afterEach(function() {
    return logout.visit();
  });

  test('it navigates to namespaces page', async function(assert) {
    assert.expect(1);
    await page.visit();
    assert.equal(
      currentRouteName(),
      'vault.cluster.access.namespaces.index',
      'navigates to the correct route'
    );
  });

  test('it should render correct number of namespaces', async function(assert) {
    assert.expect(3);
    await page.visit();
    const store = this.owner.lookup('service:store');
    // Default page size is 15
    assert.equal(store.peekAll('namespace').length, 15, 'Store has 15 namespaces records');
    assert.dom('.list-item-row').exists({ count: 15 });
    assert.dom('[data-test-list-view-pagination]').exists();
  });
});
